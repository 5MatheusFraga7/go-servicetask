# go-servicetask

- A ideia é criar uma api que leia um arquivo do tipo csv 
- Processe dados de acordo com uma regra de negócio específica
- Insira no banco os dados válidos

- [x] Estruta de dado tarefa
- [x] Criar Conexão com banco 
- [x] Criar Repositories de banco para crud de tasks
- [] Criar leitura de arquivos csv
- [] Estrutura de criação de tarefas com task-manager
- [] Criar leitura de arquivos xls


Faltou dar continuidade ao método  -> fiunalizar esse método para criar um array de tarefas vindouras de um arquivo csv

func (tm TaskManager) CreateTasks() {
	csvParser := models.CsvParser{}

	data := csvParser.GetDataFromInternalFile()

	for _, row := range data {
		deadline, _ := time.Parse("2006-01-02 15:04:05", row[1])
		tm.Tasks = append(tm.Tasks, Task{
			Name:        row[0],
			Description: row[0],
			Duration:    row[0],
			UserName:    row[0],
			Deadline:    deadline,
			Done:        false,
			FinishedAT:  nil,
		})
	}
}

