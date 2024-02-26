import psycopg2

# Parâmetros de conexão com o banco de dados
hostname = 'database-stations.cxic0so62a43.us-east-1.rds.amazonaws.com'
username = 'postgres'  # Substitua pelo seu nome de usuário
password = 'admin1234'    # Substitua pela sua senha
database = 'postgres'  # Substitua pelo nome do seu banco de dados

# Comando SQL para criar as tabelas
create_tables_sql = """
CREATE TABLE Estacao (
    Id INT PRIMARY KEY,
    latitude FLOAT,
    longitude FLOAT
);

CREATE TABLE Gas (
    Id_gas INT PRIMARY KEY,
    Id_estacao INT,
    CO2 FLOAT,
    CO FLOAT,
    NO2 FLOAT,
    MP10 FLOAT,
    MP25 FLOAT,
    FOREIGN KEY (Id_estacao) REFERENCES Estacao(Id)
);

CREATE TABLE Rad_lum (
    Id_rad INT PRIMARY KEY,
    Id_estacao INT,
    ET FLOAT,
    LI FLOAT,
    SR FLOAT,
    FOREIGN KEY (Id_estacao) REFERENCES Estacao(Id)
);
"""

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

# Executar o comando SQL para criar as tabelas
try:
    execute_sql(create_tables_sql)
    print("Tabelas criadas com sucesso!")
except Exception as e:
    print("Erro ao criar tabelas:", e)
