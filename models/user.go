package models

// The above type defines a User struct with fields for Id, Name, and Email.
// @property {int} Id - The `Id` field in the `User` struct represents the unique identifier for a
// user. It is of type `int` and is tagged with `json:"id"` for JSON marshaling purposes.
// @property {string} Name - The properties of the `User` struct are:
// @property {string} Email - The `User` struct you provided has three properties: `Id`, `Name`, and
// `Email`. The `Email` property is of type `string` and is tagged with `json:"email"`, which means it
// will be marshaled and unmarshaled as `email` when converting to
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
