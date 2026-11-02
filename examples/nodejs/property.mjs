// Viva Real Property API – Node.js example (Node 18+, native fetch, ES module)
// Usage: PILOTERR_API_KEY=your-key node property.mjs

const params = new URLSearchParams({
  "query": "https://www.vivareal.com.br/imovel/apartamento-3-quartos-brooklin-zona-sul-sao-paulo-com-garagem-186m2-venda-RS3710000-id-2776423622/",
});

const response = await fetch(`https://api.piloterr.com/v2/vivareal/property?${params}`, {
  headers: { "x-api-key": process.env.PILOTERR_API_KEY },
  signal: AbortSignal.timeout(120_000),
});

if (!response.ok) {
  console.error(`Piloterr error ${response.status}: ${await response.text()}`);
  process.exit(1);
}

console.log(JSON.stringify(await response.json(), null, 2));
