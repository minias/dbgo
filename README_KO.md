# DBGo - 동적 데이터베이스 빌더

DBGo는 JSON 입력을 기반으로 동적으로 데이터베이스 테이블을 생성할 수 있는 빌더입니다. Go로 작성되었으며 Echo 프레임워크를 사용하여 HTTP 요청을 처리합니다.

**참고: 이 README.md 파일은 OpenAI에서 개발한 ChatGPT 모델을 사용하여 생성되었습니다.**

## 사전 요구 사항

애플리케이션을 실행하기 전에 다음 사전 요구 사항이 설치되어 있는지 확인하십시오:

- Go (버전 1.16 이상)
- MySQL 또는 Oracle 데이터베이스 (선택에 따라)

## 설치

1. 저장소를 복제합니다:

```shell
git clone https://github.com/your-username/dbgo.git
```
1. 필요한 종속성을 설치합니다:

```sh
go mod download
```

## 설정

애플리케이션의 설정은 config.yml 파일에 저장됩니다. 원하는 설정으로 파일을 업데이트하세요:

- Server.Port: 서버가 실행될 포트 번호입니다.
- Server.Protocol: 서버에 사용할 프로토콜입니다. "http" 또는 "https"로 설정할 수 있습니다.
- Server.CertFile: HTTPS 프로토콜에 필요한 SSL 인증서 파일의 경로입니다.
- Server.KeyFile: HTTPS 프로토콜에 필요한 개인 키 파일의 경로입니다.
- JWTSecret: JWT 토큰 서명에 사용되는 비밀 키입니다.
- JWTExpireTime: JWT 토큰의 만료 시간(분)입니다.
- Database.Driver: 데이터베이스 드라이버 이름입니다. "mysql" 또는 "oracle" 중 하나를 선택합니다.
- Database.Source: 데이터베이스 연결 문자열입니다.

## 사용법

1. 애플리케이션을 실행합니다:

```sh
go run main.go
```

1. REST 클라이언트나 웹 브라우저를 사용하여 API 엔드포인트에 접속합니다:

- 테이블을 동적으로 생성하려면 JSON 데이터를 포함한 PUT 요청을 /table/create로 보냅니다.
- 기존 테이블을 변경하려면 JSON 데이터를 포함한 PUT 요청을 /table/alter로 보냅니다.
- 테이블을 삭제하려면 테이블 이름을 파라미터로 포함한 DELETE 요청을 /table/delete로 보냅니다.
- 사용자로 등록하려면 사용자 이메일과 비밀번호를 포함한 POST 요청을 /user/signup로 보냅니다.
- 로그아웃하려면 GET 요청을 /user/signout로 보냅니다.
- 로그인하려면 사용자 이메일과 비밀번호를 포함한 POST 요청을 /user/login로 보냅니다.
- 새 데이터베이스를 생성하려면 데이터베이스 이름을 포함한 PUT 요청을 /database/create로 보냅니다.
- 데이터베이스를 삭제하려면 데이터베이스 이름을 포함한 DELETE 요청을 /database/delete로 보냅니다.

## 라이선스
이 프로젝트는 [CC0 1.0 Universal](./LICENSE) 하에 사용됩니다.