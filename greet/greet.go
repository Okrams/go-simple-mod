package greet

// Variable de paquete se debe declarar con la palabra reservada var
var emoji = "🙋"

// ? Para exportar funciones y variables
// ? Si la primera letra esta en mayuscula esta se exporta
// ? Si la primera letra esta en minuscula no se exporta

// English retorna saludo en íngles
func English() string {
	return "Hi " + emoji
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
