const http = require('http');

const serve = http.createServer((req, res) => {
    res.setHeader( 'Content-Type', 'application/json' );

    if(req.method === 'GET' && req.url === '/health') {
        res.end('{"status":"ok"}');
        return;
    }
    if(req.method === 'POST' && req.url === '/operaciones') {
        let body = '';
        req.on('data', (chunk) => {
            body += chunk;
        });
        req.on('end', () => {
            let datos;

            try {
                datos = JSON.parse(body);
            } catch (error) {
                res.statusCode = 400;
                res.end('{"error":"Invalid JSON"}');
                return;
            }
            res.end(JSON.stringify({
                analisis: analizar(datos.Q, datos.R),
            }));
        });
        return;
    }

    res.statusCode = 404;
    res.end('{"error":"ruta no encontrada"}');

});

if(require.main === module){
serve.listen(3000, () => {
    console.log('Server is running on port 3000');
});
}

function analizar(q, r){
    const matrices = [q, r];
    let maximo = -Infinity;
    let minimo = Infinity;
    let promedio = 0;
    let suma = 0;
    let cantidad = 0;

   for(let m = 0; m < matrices.length; m++){
    const matriz = matrices[m];
    for(let i = 0; i < matriz.length; i++){
        for(let j = 0; j < matriz[i].length; j++){
            const valor = matriz[i][j];
            if(valor > maximo){
                maximo = valor;
            }
            if(valor < minimo){
                minimo = valor;
            }
            promedio += valor;
            suma += valor;
            cantidad += 1;
        }
    }
   }

const qEsDiagonal = esDiagonal(q);
const rEsDiagonal = esDiagonal(r);

   return {
    valorMaximo: maximo,
    valorMinimo: minimo,
    promedio: promedio / cantidad,
    suma: suma,
    cantidad: cantidad,
    matricesDiagonales: [
        {nombre: 'Q', esDiagonal: qEsDiagonal},
        {nombre: 'R', esDiagonal: rEsDiagonal},
    ],
   }
}

function esDiagonal(matriz){
    for(let i = 0; i < matriz.length; i++){
        for(let j = 0; j < matriz[i].length; j++){
            if(i !== j && Math.abs(matriz[i][j]) > 1e-8){
                return false;
            }
        }
    }
    return true;
}

module.exports = {
    analizar,
    esDiagonal,
};