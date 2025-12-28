# finances

Sistema de gerenciamento financeiro desenvolvido em Go.

## Estrutura do Projeto

```
finances/
├── internal/
│   └── user/           # Pacote de usuários
│       ├── entity.go
│       ├── repository.go
│       └── repository_test.go
├── migrations/         # Migrations do banco de dados
└── deployments/        # Configurações de deployment
```

## Testes

Este projeto utiliza **testes unitários** que **não acessam banco de dados real**.

### Abordagem de Testes

- Utiliza `sqlmock` para simular interações com banco de dados
- Testes unitários focam em comportamento e contratos
- Não requerem configuração de banco de dados
- Rápidos e executáveis em qualquer ambiente

### Executar Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes de um pacote específico
cd internal/user
go test -v

# Ou usar o script auxiliar
cd internal/user
./run_tests.sh
```

### Testes de Integração (Opcional)

Se precisar testar com banco de dados real:

1. **Criar banco de dados de teste:**
```bash
createdb finances_test
```

2. **Aplicar migrations:**
```bash
psql finances_test < migrations/20251227203600_add_users_table.sql
```

3. **Implementar testes de integração separados** (ver `internal/user/README_TESTS.md`)

## Dependências

- Go 1.25.3+
- PostgreSQL (apenas para produção e testes de integração)
- github.com/google/uuid
- github.com/lib/pq (driver PostgreSQL)
- github.com/DATA-DOG/go-sqlmock (testes unitários)

## Desenvolvimento

### Instalar dependências
```bash
go mod download
```

### Build
```bash
go build ./...
```

### Run tests
```bash
go test ./...
```


