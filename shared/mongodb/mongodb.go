package mongodb

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/im-jinsu/yepanmap/shared/loadconf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// mgoClient mongo client
var mgoClient *mongo.Client

// SetClient : Make MongoDB Client
func SetClient() (*mongo.Client, error) {
	var err error

	clientOpt := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s",
		loadconf.ConfigInfo.MONGODB.DBUser, loadconf.ConfigInfo.MONGODB.Password, loadconf.ConfigInfo.MONGODB.IPPort))
	if mgoClient != nil {
		if err = mgoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
			mgoClient, err = mongo.Connect(context.TODO(), clientOpt)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	} else {
		mgoClient, err = mongo.Connect(context.TODO(), clientOpt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return mgoClient, nil
}

// SetUserDBName set email to userdb name
func SetUserDBName(email string) string {
	return strings.Replace(email, ".", "#_", -1)
}
