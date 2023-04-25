package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.




import (
	"strings"
)

func ProcessString(s string) string {
	s2 := ""

	for _, letras := range s {
		if strings.ContainsAny(string(letras), "AEIOUaeiou") {
			
		}
		if strings.ContainsAny(string(letras), "bcdfghjklmnpqrstvwxyzBCDFGHKLMNPQRSTVWXYZ") {
			s2 += "." + strings.ToLower(string(letras))
		}
	}
	
	return s2
}
