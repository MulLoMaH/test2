# Правила организации подключения к базе данных в Go (на примере PostgreSQL и pgx)

## 1. Основные принципы

- **Используйте пул соединений** (`pgxpool.Pool`), а не одиночное соединение (`pgx.Conn`). Пул эффективно управляет параллельными запросами и переиспользует соединения.
- **Передавайте контекст (`context.Context`)** во все операции с БД (для таймаутов, отмены запросов, graceful shutdown).
- **Закрывайте ресурсы** (пул соединений, rows, statement) с помощью `defer`.
- **Структура репозитория должна реализовывать интерфейс** – это позволяет легко тестировать и подменять реализацию (in-memory, mock, реальная БД).
- **Инициализация подключения** – в отдельном пакете (например, `db_conn`), а слой работы с данными – в пакете `repository/postgres`.

## 2. Стандарт создания подключения + 
    package db_conn

   **Создание подключения к базе данных (пул соединений)**
    
    func NewPool(ctx context.Context) (*pgxpool.Pool, error) {                      
        connString := os.Getenv("CONN_DB")
        if connString == "" { //проверка передана ли переменная окружеения 
            return nil, errors.New("no CONN_DB in .env")
        }

   **pgxpool.New возвращает пул соединений и ошибку**
       
       pool, err := pgxpool.New(ctx, connString) 
            if err != nil {
            return nil, err
        }

   **проверяем (пингуем) соединение**

        if err := pool.Ping(ctx); err != nil { 
            return nil, err
        }
        return pool, nil
    }


## Пакет с реализацией 

    package postgres

   **Структура содержащая указатель на пул соединений**
   
    type PostgresRepo struct {
        pool *pgxpool.Pool
    }
   
   **Конструктор структуры**
   
    func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
        return &PostgresRepo{pool: pool}
    }

   **Убедимся, что реализует интерфейс**
   
    var _ repository.EmployeeRepository = (*PostgresRepo)(nil)


## Создание зависимостей

    func main() {
        ctx := context.Background()
        pool, err := db_conn.NewPool(ctx)
        if err != nil {
            log.Fatal(err)
        }
        defer pool.Close()

        repo := postgres.NewPostgresRepo(pool)   // реализация для БД
        // или repo := memory.NewEmployees()     // in-memory для тестов

        handler := handlers.New_Employee_Handlers(repo)
        // ...
    }


# Работа с Handler и HTTP 

## context 

    r.Context() — это метод структуры http.Request в Go, который возвращает контекст текущего HTTP-запроса. Этот контекст позволяет:

- Отслеживать отмену запроса (например, если клиент закрыл соединение или нажал Ctrl+C).

- Устанавливать таймауты (deadline) на выполнение обработчика.

- Передавать сквозные данные (например, ID запроса, данные аутентификации).
