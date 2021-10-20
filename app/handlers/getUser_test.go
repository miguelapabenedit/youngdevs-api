package handlers_test

// import (
// 	"fmt"
// 	"github/miguelapabenedit/youngdevs-api/app/data"
// 	"github/miguelapabenedit/youngdevs-api/app/handlers"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// type mockRepo struct{}

// func (m *mockRepo) Get(id string) *data.User {
// 	user := data.User{Email: id + "@youngdevs.com"}
// 	user.ID = 1
// 	return &user
// }
// func (m *mockRepo) Create(u *data.User) error {

// 	return nil
// }

// func TestGetUser(t *testing.T) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("/user?id=%v", 1), nil)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	handlers.NewGetUser(&mockRepo{})

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(handlers.GetUser)

// 	handler.ServeHTTP(rr, req)

// 	if status := rr.Code; status != http.StatusAccepted {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"email":"1@youngdevs.com"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
