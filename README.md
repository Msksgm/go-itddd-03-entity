# GO-ITDDD-03-ENTITY

zenn の記事「[Go でエンティティを実装（「入門ドメイン駆動設計」Chapter3）](https://zenn.dev/msksgm/articles/20220225-go-itddd-03-entity)」のサンプルコードです。

# 実行環境

- Go
  - 1.17

# 実行方法

正常形

```bash
> go run main.go -id a -name name
&{{a} name}
```

異常形

```bash
> go run main.go
2022/02/20 13:17:47 NewUserId(): userId is required
exit status 1
```

```bash
> go run main.go -id a
2022/02/20 13:18:04 user.ChangeUserName(""): name is required
exit status 1
```

```bash
> go run main.go -id a -name na
2022/02/20 13:18:31 user.ChangeUserName("na"): name na is less than three characters long
exit status 1
```

# テスト

```bash
> go test -v -count=1 ./...
?       github.com/Msksgm/go-itddd-03-entity    [no test files]
=== RUN   TestNewUser
=== RUN   TestNewUser/success
=== RUN   TestNewUser/fail_name_is_empty
=== RUN   TestNewUser/fail_name_is_less_than_three_characters
--- PASS: TestNewUser (0.00s)
    --- PASS: TestNewUser/success (0.00s)
    --- PASS: TestNewUser/fail_name_is_empty (0.00s)
    --- PASS: TestNewUser/fail_name_is_less_than_three_characters (0.00s)
=== RUN   TestChangeUserName
=== RUN   TestChangeUserName/success
=== RUN   TestChangeUserName/fail_userName_is_empty
=== RUN   TestChangeUserName/fail_userName_is_less_than_three_charcters
--- PASS: TestChangeUserName (0.00s)
    --- PASS: TestChangeUserName/success (0.00s)
    --- PASS: TestChangeUserName/fail_userName_is_empty (0.00s)
    --- PASS: TestChangeUserName/fail_userName_is_less_than_three_charcters (0.00s)
=== RUN   TestEquals
=== RUN   TestEquals/success
=== RUN   TestEquals/fail
--- PASS: TestEquals (0.00s)
    --- PASS: TestEquals/success (0.00s)
    --- PASS: TestEquals/fail (0.00s)
=== RUN   TestNewUserId
=== RUN   TestNewUserId/success
=== RUN   TestNewUserId/fail_userId_is_empty
--- PASS: TestNewUserId (0.00s)
    --- PASS: TestNewUserId/success (0.00s)
    --- PASS: TestNewUserId/fail_userId_is_empty (0.00s)
PASS
ok      github.com/Msksgm/go-itddd-03-entity/domain/model/user  0.470s
=== RUN   TestWrap
--- PASS: TestWrap (0.00s)
PASS
ok      github.com/Msksgm/go-itddd-03-entity/iterrors   0.287s
```
