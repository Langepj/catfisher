package model

type Profile struct{
  ID          bson.ObjectId `bson:"_id" json:"id"`
	DisplayName string `bson:"displayName" json:"displaame"`
	CoverImage  string `bson:"cover_image" json:"cover_image"`
	Description string `bson:"description" json:"description"`
}
