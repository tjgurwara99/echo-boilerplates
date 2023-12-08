package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/db"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/models"
)

func TestPermissionFlatten(t *testing.T) {
	// create new permissions
	permission1 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission: "permission1",
	}
	permission2 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission:        "permission2",
		ParentPermissions: []*models.Permission{&permission1},
	}
	permission3 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission:        "permission3",
		ParentPermissions: []*models.Permission{&permission2},
	}
	allPermissions := []models.Permission{permission1, permission2, permission3}

	// test flatten
	permissions, err := permission3.Flatten(allPermissions)
	require.NoError(t, err)
	require.Equal(t, []string{"permission1", "permission2", "permission3"}, permissions)
}

func TestUsersFlattenPermission(t *testing.T) {
	// create new permissions
	permission1 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission: "permission1",
	}
	permission2 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission:        "permission2",
		ParentPermissions: []*models.Permission{&permission1},
	}
	permission3 := models.Permission{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permission:        "permission3",
		ParentPermissions: []*models.Permission{&permission2},
	}
	require.NoError(t, db.DB.Create(&permission1).Error)
	require.NoError(t, db.DB.Create(&permission2).Error)
	require.NoError(t, db.DB.Create(&permission3).Error)

	// create new user
	user := models.User{
		Model: models.Model{
			ID: uuid.New(),
		},
		Permissions: []*models.Permission{&permission3},
	}

	// test flatten
	permissions, err := user.FlatPermissions(db.DB)
	require.NoError(t, err)
	require.Equal(t, []string{"permission1", "permission2", "permission3"}, permissions)
}
