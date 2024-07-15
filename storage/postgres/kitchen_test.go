package postgres

import (
	pb "auth_service/genproto/kitchen"
	"context"
	"fmt"
	"testing"
)

func newKitchenRepo() *KitchenRepo {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	return &KitchenRepo{Db: db}
}

// func TestCreateKitchen(t *testing.T) {
// 	k := newKitchenRepo()

// 	kitchen := pb.ReqCreateKitchen{
// 		OwnerId:     "2b2ce16f-ed08-41cf-b26c-42facc0f617c",
// 		Name:        "Local",
// 		Description: "good",
// 		CuisineType: "home made",
// 		Address:     "st 6 avenue",
// 		PhoneNumber: "+9985525555",
// 	}
// 	_, err := k.CreateKitchen(context.Background(), &kitchen)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestUpdateKitchen(t *testing.T) {
// 	k := newKitchenRepo()

// 	kitchen := pb.ReqUpdateKitchen{
// 		Id:          "39f996de-1607-4dab-8c38-733532887af6",
// 		OwnerId:     "2b2ce16f-ed08-41cf-b26c-42facc0f617c",
// 		Name:        "rtgghjkfgbn b",
// 		Description: "tgvbnmb",
// 		CuisineType: "tgfghjkb",
// 		Address:     "tyhnghjk",
// 		PhoneNumber: "6tyghjkhn",
// 	}
// 	_, err := k.UpdateKitchen(context.Background(), &kitchen)
// 	if err != nil {
// 		t.Fatalf("UpdateKitchen failed: %v", err)
// 	}
// }

// func TestGetKitchenById(t *testing.T) {
// 	k := newKitchenRepo()

// 	ctx := context.Background()

// 	kitchenID := "39f996de-1607-4dab-8c38-733532887af6"

// 	kitchen, err := k.GetKitchenById(ctx, kitchenID)
// 	if err != nil {
// 		t.Fatalf("GetKitchenById failed: %v", err)
// 	}

// 	if kitchen == nil {
// 		t.Fatalf("Expected kitchen to be found, but got nil")
// 	}
// 	if kitchen.Id != kitchenID {
// 		t.Errorf("Expected kitchen ID %s, but got %s", kitchenID, kitchen.Id)
// 	}
// }

// func TestGetKitchens(t *testing.T) {
// 	k := newKitchenRepo()

// 	ctx := context.Background()

// 	filter := &pb.Pagination{
// 		Page:  1,
// 		Limit: 10,
// 	}

// 	kitchens, err := k.GetKitchens(ctx, filter)
// 	if err != nil {
// 		t.Fatalf("GetKitchens failed: %v", err)
// 	}

// 	if kitchens == nil {
// 		t.Fatalf("Expected kitchens list to be returned, but got nil")
// 	}
// 	if len(kitchens.Kitchens) == 0 {
// 		t.Fatalf("Expected non-empty kitchens list, but got empty")
// 	}
// 	for _, kitchen := range kitchens.Kitchens {
// 		if kitchen == nil {
// 			t.Fatalf("Found nil kitchen in the list")
// 		}
// 	}
// }

func TestSearchKitchens(t *testing.T) {
	k := newKitchenRepo()

	ctx := context.Background()

	filter := &pb.Search{
		Name:        "Local",
		CuisineType: "home made",
		Rating:      0,
		Address:     "",
		Page:        1,
		Limit:       10,
	}
	// "Local home made"

	kitchens, err := k.SearchKitchens(ctx, filter)
	if err != nil {
		t.Fatalf("SearchKitchens failed: %v", err)
	}
	fmt.Println(kitchens)
	if kitchens == nil {
		t.Fatalf("Expected kitchens list to be returned, but got nil")
	}
	if len(kitchens.Kitchens) == 0 {
		t.Fatalf("Expected non-empty kitchens list, but got empty")
	}
	for _, kitchen := range kitchens.Kitchens {
		if kitchen == nil {
			t.Fatalf("Found nil kitchen in the list")
		}
	}
}
