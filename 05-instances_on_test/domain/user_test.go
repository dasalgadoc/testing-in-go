package domain

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Traditional instance
func TestUser_UserCanEditBlogWithEnough(t *testing.T) {
	user := traditionalInstance("John Doe", 2)
	assert.True(t, user.CanEditBlog())
}

func TestUser_UserCanEditBlogWithNotEnough(t *testing.T) {
	user := traditionalInstance("John Doe", 1)
	assert.False(t, user.CanEditBlog())
}

// Builder instance: good for semantics but no for consistency
func TestUser_UserCanEditBlogWithEnoughBuilder(t *testing.T) {
	user := NewUserBuilder().
		WithAccessLevel(2).
		Build()

	assert.True(t, user.CanEditBlog())
}

func TestUser_UserCanEditBlogWithNotEnoughBuilder(t *testing.T) {
	user := NewUserBuilder().
		WithAccessLevel(1).
		Build()

	assert.False(t, user.CanEditBlog())
}

// Object Mother instance: with stubs, good for randomization and consistency but no for composition
func TestUser_UserCanEditBlogWithEnoughObjectMother(t *testing.T) {
	user := UserWithAccessLevel(2)

	assert.True(t, user.CanEditBlog())
	assert.NotEmpty(t, user.Name())
	assert.NotEmpty(t, user.ID())
}

func TestUser_UserCanEditBlogWithNotEnoughObjectMother(t *testing.T) {
	user := UserWithAccessLevel(1)

	assert.False(t, user.CanEditBlog())
	assert.NotEmpty(t, user.Name())
	assert.NotEmpty(t, user.ID())
}

// Object Mother with stub that compound the object: If the language allows it, use named parameters/arguments
func TestUser_UserCanEditBlogWithEnoughObjectMotherWithStub(t *testing.T) {
	testData := map[string]any{
		"name":        "John Doe",
		"accessLevel": 2,
	}
	user := UserWithParameters(testData)

	assert.True(t, user.CanEditBlog())
	assert.NotEmpty(t, user.Name())
	assert.NotEmpty(t, user.ID())
}

func TestUser_UserCanEditBlogWithNotEnoughObjectMotherWithStub(t *testing.T) {
	testData := map[string]any{
		"name":        "John Doe",
		"accessLevel": 1,
	}
	user := UserWithParameters(testData)

	assert.False(t, user.CanEditBlog())
	assert.NotEmpty(t, user.Name())
	assert.NotEmpty(t, user.ID())
}

func traditionalInstance(name string, accessLevel int) User {
	userName, err := NewUserName(name)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	userAccessLevel, err := NewUserAccessLevelFromInt(accessLevel)
	if err != nil {
		log.Panic(err)
		panic(err)
	}
	user, err := NewUser(*userName, *userAccessLevel)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	return *user
}
