package gqlgen_todos


type Todo {
	id: ID!
	text: String!
	done: Boolean!
	user: User!
}