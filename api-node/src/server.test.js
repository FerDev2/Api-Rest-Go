const assert = require('assert');
const { analizar } = require('./server');

const resultado = analizar(
  [[0, 1], [1, 0]],
  [[1, 0], [0, 1]]
);

assert.strictEqual(resultado.valorMaximo, 1);
assert.strictEqual(resultado.valorMinimo, 0);
assert.strictEqual(resultado.suma, 4);
assert.strictEqual(resultado.cantidad, 8);
assert.strictEqual(resultado.promedio, 0.5);
assert.strictEqual(resultado.matricesDiagonales[0].esDiagonal, false);
assert.strictEqual(resultado.matricesDiagonales[1].esDiagonal, true);

console.log('ok');