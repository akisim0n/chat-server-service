-- +goose Up
-- +goose StatementBegin
INSERT INTO CHAT
values (
    1,
    'UsersChat',
    'Chat for all!'
    ),
    (
        2,
        'TestUsersChat',
        'Chat for test users'
    ),
   (
       3,
       'PrivateChat',
       'Do not tell about Private Chat'
   )
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM CHAT
WHERE ID IN (1,2,3)
-- +goose StatementEnd
