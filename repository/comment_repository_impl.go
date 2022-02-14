package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"gomysql/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = id
	return comment, nil

}

func (repo *commentRepositoryImpl) Update(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "UPDATE comments SET email = ?, comment = ? WHERE id = ?"
	row, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment, comment.Id)
	if err != nil {
		return entity.Comment{}, err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return entity.Comment{}, err
	}

	if count == 0 {
		return entity.Comment{}, errors.New("Id " + strconv.Itoa(int(comment.Id)) + " not found")
	}

	return comment, nil
}

func (repo *commentRepositoryImpl) Delete(ctx context.Context, id int64) (entity.Comment, error) {
	sqlExec := "DELETE FROM comments WHERE id = ?"
	row, err := repo.DB.ExecContext(ctx, sqlExec, id)
	if err != nil {
		return entity.Comment{}, err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return entity.Comment{}, err
	}

	if count == 0 {
		return entity.Comment{}, errors.New("Data Not Found")
	}

	return entity.Comment{}, nil

}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	sqlExec := "SELECT * FROM comments WHERE id = ?"
	row, err := repo.DB.QueryContext(ctx, sqlExec, id)
	if err != nil {
		return entity.Comment{}, err
	}
	defer row.Close()
	comment := entity.Comment{}
	// for row.Next() {
	// 	err := row.Scan(&comment.Id, &comment.Email, &comment.Comment)
	// 	if err != nil {
	// 		return entity.Comment{}, err
	// 	}
	// }
	// return comment, nil
	if row.Next() {
		row.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sqlExec := "SELECT * FROM comments"
	rows, err := repo.DB.QueryContext(ctx, sqlExec)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
