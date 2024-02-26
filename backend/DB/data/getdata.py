import psycopg2

# Parâmetros de conexão com o banco de dados
hostname = 'database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com'
username = 'postgres'  # Substitua pelo seu nome de usuário
password = 'admin1234'    # Substitua pela sua senha
database = 'postgres'  # Substitua pelo nome do seu banco de dados

# Função para executar consultas SQL SELECT
def select_data(table_name):
    select_sql = f"SELECT * FROM {table_name}"
    try:
        conn = psycopg2.connect(
            host=hostname,
            user=username,
            password=password,
            database=database
        )
        cur = conn.cursor()
        cur.execute(select_sql)
        rows = cur.fetchall()
        print(f"Dados da tabela {table_name}:")
        for row in rows:
            print(row)
        conn.close()
    except Exception as e:
        print(f"Erro ao selecionar dados da tabela {table_name}:", e)

# Ver dados nas tabelas
select_data("Estacao")
select_data("Gas")
select_data("Rad_lum")