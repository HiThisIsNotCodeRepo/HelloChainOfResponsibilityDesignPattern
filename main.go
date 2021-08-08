package main

func main() {
	aCashier := &cashier{}
	aMedical := &medical{}
	aMedical.setNext(aCashier)
	aDoctor := &doctor{}
	aDoctor.setNext(aMedical)
	aReception := &reception{}
	aReception.setNext(aDoctor)
	aPatient := &patient{name: "Patient"}
	aReception.execute(aPatient)
	aReception.execute(aPatient)
}
