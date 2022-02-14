package repository

import (
	"context"
	"fmt"
	"gomysql"
	"gomysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())
	comment := entity.Comment{
		Email:   "new@anakdesa.id",
		Comment: "test comment",
	}

	result, err := commentRepository.Insert(context.Background(), comment)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)
}

func TestCommentUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())
	comment := entity.Comment{
		Id:      6,
		Email:   "bowo@anakdesa.id",
		Comment: "test comment update",
	}

	result, err := commentRepository.Update(context.Background(), comment)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result)

}

func TestCommentDelete(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	comment := entity.Comment{
		Id: 4,
	}

	_, err := commentRepository.Delete(context.Background(), comment.Id)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Delete Success Id", comment.Id)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 200)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(comment)

}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(gomysql.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
