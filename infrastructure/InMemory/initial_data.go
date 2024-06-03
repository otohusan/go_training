package inmemory

import (
	"go-training/domain/model"
	"time"
)

var Users = []*model.User{
	{
		ID:        "a",
		Name:      "a",
		Password:  "a",
		Email:     "a",
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

var StudySets = []*model.StudySet{
	{
		ID:          "1",
		UserID:      "a",
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

var Flashcards = []*model.Flashcard{
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

var Favorites = []*model.Favorite{
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

var EmailVerification = []*model.EmailVerification{}
