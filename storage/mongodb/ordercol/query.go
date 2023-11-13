package ordercol

import (
	"context"
	"fmt"
	"github.com/ponlv/go-kit/mongodb"
	bsonutil "github.com/ponlv/go-kit/mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// create
func Create(ctx context.Context, data *Order) (interface{}, error) {
	// get collection
	coll := mongodb.Coll(os.Getenv("MONGODB_DATABASE"), data)

	// create
	id, err := coll.CreateWithCtx(ctx, data)
	if err != nil {
		return nil, err
	}

	// end
	return id, nil
}

// update
func Update(ctx context.Context, data *Order) (bool, error) {

	objID, err := primitive.ObjectIDFromHex(data.GetIDString())
	if err != nil {
		return false, err
	}
	// filter by project id
	filter := bsonutil.BsonAdd(nil, "_id", objID)
	// update
	update := bsonutil.BsonSetMap(nil,
		bsonutil.ConvertStructToBSONMap(
			data,
			&bsonutil.MappingOpts{RemoveID: true},
		),
	)
	// update
	collection := mongodb.Coll(os.Getenv("MONGODB_DATABASE"), &Order{})
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func FindWithID(ctx context.Context, id string) (*Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	// generate filter
	filter := bsonutil.BsonAdd(nil, "_id", objID)
	// end
	return findWithCondition(ctx, filter)
}

func FindAllEmptyVnPostID(ctx context.Context) ([]*Order, error) {
	coll := mongodb.CollRead(os.Getenv("MONGODB_DATABASE"), &Order{})

	option := new(options.FindOptions)
	option.SetSort(bson.D{{Key: "created_at", Value: -1}})

	filter := bson.D{}

	filter = bsonutil.BsonAdd(filter, "vn_post_id", "")
	filter = bsonutil.BsonAdd(filter, "carrier_id", 12)
	filter = bsonutil.BsonNotIn(filter, "depot_id", []int{
		138773, // thai
		143753, // us
	})

	var orders []*Order
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	} else {
		if err = cursor.All(ctx, &orders); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func FindAllOrderHaveVnPostID(ctx context.Context) ([]*Order, error) {
	coll := mongodb.CollRead(os.Getenv("MONGODB_DATABASE"), &Order{})

	option := new(options.FindOptions)
	option.SetSort(bson.D{{Key: "created_at", Value: -1}})

	filter := bson.D{}

	filter = bsonutil.BsonNotEqual(filter, "vn_post_id", "")
	filter = bsonutil.BsonNotIn(filter, "status", []string{
		"Success",
	})

	var orders []*Order
	cursor, err := coll.Find(ctx, filter, option)
	if err != nil {
		return nil, err
	} else {
		if err = cursor.All(ctx, &orders); err != nil {
			return nil, err
		}
	}

	return orders, nil
}

func FindManyWithFilter(ctx context.Context, page int64, limit int64, search string, status []string) (*Pagination, error) {
	coll := mongodb.CollRead(os.Getenv("MONGODB_DATABASE"), &Order{})
	//init option
	options := new(options.FindOptions)
	options.SetSkip(int64((page - 1) * limit))
	options.SetLimit(int64(limit))
	options.SetSort(bson.D{{Key: "created_at", Value: -1}})
	//
	//Init Filter
	//for search
	filter := bson.D{}
	//
	//By: user's info
	//
	if search != "" {
		filter = append(filter, bson.E{
			Key: "$or",
			Value: []bson.M{
				{"first_name": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"last_name": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"email": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"apartment": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"city": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"country": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
				{"order_id": bson.M{
					"$regex":   search,
					"$options": "im",
				}},
			}})
	}

	//
	//By: status
	if len(status) > 0 {
		filter = bsonutil.BsonIn(filter, "status", status)
	}

	total, _ := coll.CountDocuments(ctx, filter)
	var orders []*Order
	cursor, err := coll.Find(ctx, filter, options)
	fmt.Println(err)
	if err != nil {
		return nil, err
	} else {
		if err = cursor.All(ctx, &orders); err != nil {
			return nil, err
		}
	}
	return &Pagination{
		Orders: orders,
		Total:  total,
	}, nil

}

// find common
func findWithCondition(ctx context.Context, filter interface{}, findOptions ...*options.FindOneOptions) (*Order, error) {
	coll := mongodb.CollRead(os.Getenv("MONGODB_DATABASE"), &Order{})
	result := &Order{}
	if err := coll.FirstWithCtx(ctx, filter, result, findOptions...); err != nil {
		return nil, err
	}

	return result, nil
}
