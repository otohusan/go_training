package inmemory

import (
	"go-training/domain/model"
	"time"
)

func InitializeUsers() []*model.User {
	return []*model.User{
		{
			ID:        "1",
			Name:      "user1",
			Password:  "password1",
			Email:     "sasasa",
			CreatedAt: time.Now(),
		},
		{
			ID:        "2",
			Name:      "user2",
			Password:  "password2",
			Email:     "user2@example.com",
			CreatedAt: time.Now(),
		},
	}
}

func InitializeStudySets() []*model.StudySet {
	return []*model.StudySet{
		{
			ID:          "1",
			UserID:      "1",
			Title:       "Study Set 1",
			Description: "Description 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "2",
			UserID:      "2",
			Title:       "Study Set 2",
			Description: "Description 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}

func InitializeFlashcards() []*model.Flashcard {
	return []*model.Flashcard{
		{
			ID:         "1",
			StudySetID: "1",
			Question:   "Question 1",
			Answer:     "Answer 1",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         "2",
			StudySetID: "2",
			Question:   "Question 2",
			Answer:     "Answer 2",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}
}

func InitializeFavorites() []*model.Favorite {
	favorites := []*model.Favorite{
		{
			ID:         "1",
			UserID:     "1",
			StudySetID: "1",
			CreatedAt:  time.Now(),
		},
		{
			ID:         "2",
			UserID:     "1",
			StudySetID: "2",
			CreatedAt:  time.Now(),
		},
	}
	return favorites
}
