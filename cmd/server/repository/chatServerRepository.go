package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	chat "github.com/akisim0n/chat-server-service/pkg/chatServer_v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatServerRepository struct {
	DB *pgxpool.Pool
	chat.UnimplementedChatServerV1Server
}

func NewChatServerRepository(db *pgxpool.Pool) *ChatServerRepository {
	return &ChatServerRepository{
		DB: db,
	}
}

func (repo *ChatServerRepository) Create(ctx context.Context, request *chat.CreateRequest) (*chat.CreateResponse, error) {

	builderInsert := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("title").
		Values(request.GetTitle()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, errorHandler("failed to create query: %v", err)
	}

	var newChatId int64

	err = repo.DB.QueryRow(ctx, query, args...).Scan(&newChatId)
	if err != nil {
		return nil, errorHandler("failed to insert: %v", err)
	}

	for _, userId := range request.GetUserIds() {
		insertUserQuery := sq.Insert("links").
			PlaceholderFormat(sq.Dollar).
			Columns("chat_id", "user_id").
			Values(newChatId, userId)

		query, args, err = insertUserQuery.ToSql()
		if err != nil {
			return nil, errorHandler("failed to create query: %v", err)
		}

		_, err = repo.DB.Exec(ctx, query, args...)
		if err != nil {
			return nil, errorHandler("failed to insert in links: %v", err)
		}
	}

	return &chat.CreateResponse{Id: newChatId}, nil
}

func (repo *ChatServerRepository) Delete(ctx context.Context, request *chat.DeleteRequest) (*emptypb.Empty, error) {
	// Delete from "links"
	deleteBuilder := sq.Delete("links").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"chat_id": request.GetId()})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to create query: %v", err)
	}

	_, err = repo.DB.Exec(ctx, query, args...)
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to delete links: %v", err)
	}
	// Delete from "messages"
	deleteBuilder = sq.Delete("messages").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"chat_id": request.GetId()})

	query, args, err = deleteBuilder.ToSql()
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to create query: %v", err)
	}

	_, err = repo.DB.Exec(ctx, query, args...)
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to delete messages: %v", err)
	}

	// Delete "chat"
	deleteBuilder = sq.Delete("chat").PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": request.GetId()})

	query, args, err = deleteBuilder.ToSql()
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to create query: %v", err)
	}

	_, err = repo.DB.Exec(ctx, query, args...)
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to delete chat: %v", err)
	}

	return new(emptypb.Empty), nil
}

func (repo *ChatServerRepository) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*emptypb.Empty, error) {

	insertBuilder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text").
		Values(request.GetChatId(), request.GetUserId(), request.GetText())

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to create query: %v", err)
	}

	_, err = repo.DB.Exec(ctx, query, args...)
	if err != nil {
		return new(emptypb.Empty), errorHandler("failed to insert: %v", err)
	}

	return new(emptypb.Empty), nil
}

func errorHandler(text string, err error) error {
	return errors.New(fmt.Sprintf(text, err))
}
