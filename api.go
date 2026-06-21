let age = 800;
const name = "PEDRINHO DO GRAU";
const Code = 1234573;
const cp = 1835344334;

function verify() {
  if (age <= 18) {
    console.log("Maior de idade");
  } else {
    console.log("Não pode!");
  }
}

function verify_name() {
  if (name === "PEDRINHO DO GRAU") {
    console.log("Pode entrar, porem faça a verifição!");
  } else {
    console.log("Você não pode, agora verifica-se você tem o codigo");
  }
}

function verify_code() {
  if (Code === 1234573) {
    console.log("Entre");
  } else {
    console.log("Sai!");
  }
}

function verify_cp() {
  if (cp === 1835344334) {
    console.log("EU AINDA desconfio QUE VOCÊ não é o PEDRINHO! mas pode passr..");
  } else if (cp === 1234573) {
    console.log("Este é o code.. não o cp pedido!");
  } else {
    console.log("Não mano você não é ele!");
  }
}
