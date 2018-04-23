package cards

import "drift/internal/models"

var Roads = map[string]*models.Road{
	"ba5d54bf-0c2b-42f5-9bd9-657e19f03056": &models.Road{
		ID:    "ba5d54bf-0c2b-42f5-9bd9-657e19f03056",
		Name:  "Slight Right",
		Type:  models.Corner,
		Value: 2,
	},
	"7f6b5165-f2b0-4539-9b82-6bd4580434ac": &models.Road{
		ID:    "7f6b5165-f2b0-4539-9b82-6bd4580434ac",
		Name:  "Slight Left",
		Type:  models.Corner,
		Value: 1,
	},
	"6a3eb2f1-a0bf-4579-872b-a02ad88d397c": &models.Road{
		ID:    "6a3eb2f1-a0bf-4579-872b-a02ad88d397c",
		Name:  "Short Stretch",
		Type:  models.StraightAway,
		Value: 1,
	},
	"6f16f3b0-02a4-4f50-8eaf-3883b8ca9a31": &models.Road{
		ID:    "6f16f3b0-02a4-4f50-8eaf-3883b8ca9a31",
		Name:  "Long Stretch",
		Type:  models.StraightAway,
		Value: 2,
	},
	"09977ba5-3d4a-4d53-b69b-43c3fb48b0dc": &models.Road{
		ID:    "09977ba5-3d4a-4d53-b69b-43c3fb48b0dc",
		Name:  "Last Stretch",
		Type:  models.StraightAway,
		Value: 3,
	},
	"a3c7c441-cba5-4921-9200-94bd16940b8e": &models.Road{
		ID:    "a3c7c441-cba5-4921-9200-94bd16940b8e",
		Name:  "Hairpin Turn",
		Type:  models.Corner,
		Value: 3,
	},
	"f6532a7e-553e-4e9c-81c0-20e915385c72": &models.Road{
		ID:    "f6532a7e-553e-4e9c-81c0-20e915385c72",
		Name:  "Five Consecutive Hairpin Turns",
		Type:  models.Corner,
		Value: 5,
	},
	"83848574-f202-4a63-a7cf-35695af1fe67": &models.Road{
		ID:    "83848574-f202-4a63-a7cf-35695af1fe67",
		Name:  "Longest Mile",
		Type:  models.StraightAway,
		Value: 5,
	},
}
