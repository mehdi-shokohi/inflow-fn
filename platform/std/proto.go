package std

type ProtocolHeaderV1[T any] struct{
	Headers T `json:"_headers" bson:"headers"`
}

type ProtocolBodyV1[T any] struct{
	Body T `json:"_data" bson:"data"`
}