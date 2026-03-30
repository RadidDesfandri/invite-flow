"use client";

import { useState } from "react";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";

type PingResponse = {
  success: boolean;
  message: string;
  data: {
    message: string;
    time: string;
  };
};

export default function Home() {
  const [result, setResult] = useState<string>(
    "Click the button to call backend",
  );
  const [loading, setLoading] = useState(false);

  const pingBackend = async () => {
    try {
      setLoading(true);
      const { data } = await api.get<PingResponse>("/api/v1/ping");
      setResult(`${data.data.message} at ${data.data.time}`);
    } catch {
      setResult("Failed to reach backend API");
    } finally {
      setLoading(false);
    }
  };

  return (
    <main className="mx-auto flex min-h-screen w-full max-w-3xl flex-col items-center justify-center gap-6 p-6">
      <h1 className="text-3xl font-bold">Invite Flow Starter</h1>
      <p className="text-muted-foreground text-center">
        Next.js + Tailwind + shadcn + Axios connected to Go backend.
      </p>
      <Button size="lg" onClick={pingBackend} disabled={loading}>
        {loading ? "Loading..." : "Ping Backend"}
      </Button>
      <code className="rounded-md border px-3 py-2 text-sm">{result}</code>
    </main>
  );
}
