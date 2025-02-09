// services/ranking_service.go
package services

import (
	"backend/models"
	"backend/repositories"
	"time"
)

type RankingService interface {
    CalculateRanking(clubID uint) (*models.Ranking, error)
    UpdateRanking(clubID uint) (*models.Ranking, error)
    GetAllRankings() ([]models.ClubRanking, error)
    UpdateCriteria(criteria *models.RankingCriteria) error
    GetCriteria() ([]models.RankingCriteria, error)
}

type rankingService struct {
    rankingRepo repositories.RankingRepository
    clubRepo    repositories.ClubRepository
    domainRepo  repositories.DomainRepository
    studentRepo repositories.StudentRepository
    scoreRepo   repositories.ScoreRepository
}
func NewRankingService(
    rankingRepo repositories.RankingRepository,
    clubRepo repositories.ClubRepository,
    domainRepo repositories.DomainRepository,
    studentRepo repositories.StudentRepository,
) RankingService {
    return &rankingService{
        rankingRepo: rankingRepo,
        clubRepo:    clubRepo,
        domainRepo:  domainRepo,
        studentRepo: studentRepo,
    }
}

func (s *rankingService) CalculateRanking(clubID uint) (*models.Ranking, error) {
    domains, err := s.domainRepo.GetByClubID(clubID)
    if err != nil {
        return nil, err
    }

    criteria, err := s.rankingRepo.GetCriteria()
    if err != nil {
        return nil, err
    }

    var totalScore float64
    for _, domain := range domains {
        for _, criterion := range criteria {
            score := s.calculateDomainScore(domain, criterion)
            totalScore += score * criterion.Weight
        }
    }

    ranking := &models.Ranking{
        ClubID: clubID,
        Score:  totalScore,
    }

    err = s.rankingRepo.Create(ranking)
    return ranking, err
}

func (s *rankingService) UpdateRanking(clubID uint) (*models.Ranking, error) {
    ranking, err := s.CalculateRanking(clubID)
    if err != nil {
        return nil, err
    }
    
    // Handle error from Update
    if err := s.rankingRepo.Update(ranking); err != nil {
        return nil, err
    }
    
    return ranking, nil
}

func (s *rankingService) GetAllRankings() ([]models.ClubRanking, error) {
    rankings, err := s.rankingRepo.GetAll()
    if err != nil {
        return nil, err
    }

    var clubRankings []models.ClubRanking
    for i, ranking := range rankings {
        club, err := s.clubRepo.GetClubByID(ranking.ClubID)
        if err != nil {
            continue
        }
        clubRankings = append(clubRankings, models.ClubRanking{
            ClubID:       club.ID,
            ClubName:     club.Name,
            OverallScore: ranking.Score,
            Rank:         i + 1,
        })
    }
    return clubRankings, nil
}

func (s *rankingService) UpdateCriteria(criteria *models.RankingCriteria) error {
    return s.rankingRepo.UpdateCriteria(criteria)
}

func (s *rankingService) GetCriteria() ([]models.RankingCriteria, error) {
    return s.rankingRepo.GetCriteria()
}

// services/ranking_service.go

func (s *rankingService) calculateDomainScore(domain models.Domain, criteria models.RankingCriteria) float64 {
    const (
        RECENT_MONTHS = 3
        MIN_SCORES = 3
    )

    // Get all students in the domain
    students, err := s.studentRepo.GetStudentsByDomainID(domain.ID)
    if err != nil {
        return 0.0
    }

    var domainTotalScore float64
    var activeStudents int

    for _, student := range students {
        // Calculate student's score
        studentScore, hasMinScores := s.calculateStudentScore(student.ID, RECENT_MONTHS, MIN_SCORES)
        if hasMinScores {
            domainTotalScore += studentScore * criteria.Weight
            activeStudents++
        }
    }

    // Return average domain score
    if activeStudents == 0 {
        return 0.0
    }
    return domainTotalScore / float64(activeStudents)
}

func (s *rankingService) calculateStudentScore(studentID uint, recentMonths, minScores int) (float64, bool) {
    const (
        RECENT_WEIGHT = 0.7
        HISTORICAL_WEIGHT = 0.3
    )

    now := time.Now()
    cutoffDate := now.AddDate(0, -recentMonths, 0)

    // Get student's scores
    scores, err := s.scoreRepo.GetStudentScores(studentID)
    if err != nil {
        return 0.0, false
    }

    var recentScores, historicalScores []float64

    for _, score := range scores {
        if score.Month.After(cutoffDate) {
            recentScores = append(recentScores, float64(score.Score))
        } else {
            historicalScores = append(historicalScores, float64(score.Score))
        }
    }

    // Check if student has minimum required scores
    totalScores := len(recentScores) + len(historicalScores)
    if totalScores < minScores {
        return 0.0, false
    }

    // Calculate weighted average
    recentAvg := calculateAverage(recentScores)
    historicalAvg := calculateAverage(historicalScores)

    finalScore := (recentAvg * RECENT_WEIGHT) + (historicalAvg * HISTORICAL_WEIGHT)
    return finalScore, true
}

func calculateAverage(scores []float64) float64 {
    if len(scores) == 0 {
        return 0.0
    }
    
    var sum float64
    for _, score := range scores {
        sum += score
    }
    return sum / float64(len(scores))
}