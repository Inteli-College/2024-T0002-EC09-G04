import psycopg2
import psycopg2.extras

# Parâmetros de conexão com o banco de dados
hostname = 'database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com'
username = 'postgres'  # Substitua pelo seu nome de usuário
password = 'admin1234'    # Substitua pela sua senha
database = 'postgres'  # Substitua pelo nome do seu banco de dados

# Dados fictícios para inserção nas tabelas
estacoes_data = [
    (1, 36.000000, 16.000000)
]

gas_data = [
    (1, 1, 526.000000, 7.000000, 343.000000, 404.000000, 179.000000)
]

rad_lum_data = [
    (1, 1, 267.000000, 3.000000, 241.000000)
]

# Função para executar comandos SQL
def execute_sql(sql):
    conn = psycopg2.connect(
        host=hostname,
        user=username,
        password=password,
        database=database
    )
    cur = conn.cursor()
    cur.execute(sql)
    conn.commit()
    conn.close()

# Função para inserir dados nas tabelas
def insert_data(table_name, data):
    insert_sql = f"INSERT INTO {table_name} VALUES %s"
    try:
        conn = psycopg2.connect(
            host=hostname,
            user=username,
            password=password,
            database=database
        )
        cur = conn.cursor()
        psycopg2.extras.execute_values(cur, insert_sql, data)
        conn.commit()
        conn.close()
        print(f"Dados inseridos na tabela {table_name} com sucesso!")
    except Exception as e:
        print(f"Erro ao inserir dados na tabela {table_name}:", e)

# Inserir dados nas tabelas
insert_data("Estacao", estacoes_data)
insert_data("Gas", gas_data)
insert_data("Rad_lum", rad_lum_data)
