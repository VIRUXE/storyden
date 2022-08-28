/*
 * storyden
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ThreadSubmission struct {
	// The title of the thread.
	Title string `json:"title"`
	// The markdown body for the new thread.
	Body string `json:"body"`
	// A list of tags for the new thread.
	Tags []string `json:"tags"`

	Category string `json:"category"`
}