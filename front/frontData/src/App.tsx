import { useState } from "react";
import "./App.css";
import type { Datos } from "./types/datos.type";

function App() {
  const [analisis, setAnalisis] = useState<Datos | null>(null);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError(null);

    const loginResponse = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        usuario: import.meta.env.VITE_API_USER,
        contrasena: import.meta.env.VITE_API_PASSWORD,
      }),
    });
    const sesion = await loginResponse.json();
    if (!loginResponse.ok) {
      setError(sesion.error ?? "No se pudo iniciar sesión");
      return;
    }

    const formData = new FormData(e.target as HTMLFormElement);
    const numeros = Array.from(formData.values()).map((value) =>
      parseInt(value as string),
    );

    const matriz = [numeros.slice(0, 3), numeros.slice(0, 3)];
    const response = await fetch("http://localhost:8080/qr", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${sesion.token}`,
      },
      body: JSON.stringify({ matriz }),
    });
    const data = await response.json();
    if (!response.ok) {
      setError(data.error ?? "No se pudo analizar la matriz");
      return;
    }
    setAnalisis(data.analisis);
  };
  return (
    <section
      style={{
        display: "flex",
        flexDirection: "column",
        gap: "10px",
        alignItems: "center",
        justifyContent: "center",
        width: "80%",
        margin: "0 auto",
      }}
    >
      <h1>Análisis de Matrices</h1>

      <form
        onSubmit={handleSubmit}
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "10px",
          width: "100%",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <div style={{ display: "flex", gap: "10px" }}>
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz1"
            placeholder="Ingrese la matriz"
          />
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz2"
            placeholder="Ingrese la matriz"
          />
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz3"
            placeholder="Ingrese la matriz"
          />
        </div>

        <div style={{ display: "flex", gap: "10px" }}>
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz4"
            placeholder="Ingrese la matriz"
          />
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz5"
            placeholder="Ingrese la matriz"
          />
          <input
            type="number"
            max={9}
            min={0}
            style={{ width: "25px", height: "25px" }}
            onInput={(e) => {
              if ((e.target as HTMLInputElement).value.length > 1) {
                (e.target as HTMLInputElement).value = (
                  e.target as HTMLInputElement
                ).value.slice(0, 1);
              }
            }}
            name="matriz6"
            placeholder="Ingrese la matriz"
          />
        </div>
        <button type="submit">Analizar</button>
        {error && <p>{error}</p>}
      </form>

      {analisis && (
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            gap: "10px",
            alignItems: "start",
            justifyContent: "start",
          }}
        >
          <h2>Análisis</h2>
          <p>{JSON.stringify(analisis, null, 2)}</p>

          <h2>cantidad</h2>
          <p>{analisis.cantidad}</p>
          <h2>matricesDiagonales</h2>
          <p>{JSON.stringify(analisis.matricesDiagonales, null, 2)}</p>
          <h2>promedio</h2>
          <p>{analisis.promedio}</p>
          <h2>suma</h2>
          <p>{analisis.suma}</p>
          <h2>valorMaximo</h2>
          <p>{analisis.valorMaximo}</p>
          <h2>valorMinimo</h2>
        </div>
      )}
    </section>
  );
}

export default App;
