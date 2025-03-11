import { Card } from "@/components/ui/card";
import { LinkIcon } from "lucide-react";
import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="font-sans antialiased">
        <main className="min-h-svh bg-gradient-to-b from-blue-100 to-white flex flex-col items-center justify-center p-4">
          <Card className="w-full max-w-2xl p-8 space-y-8">
            {/* card header */}
            <div className="text-center space-y-2">
              <LinkIcon className="h-12 w-12 text-blue-500 mx-auto" />
              <h1 className="text-3xl font-bold text-gray-800">
                URL Shortener
              </h1>
              <p className="text-gray-600">
                Simplify your links with our powerful URL shortening service
              </p>
            </div>

            {children}
          </Card>
        </main>
      </body>
    </html>
  );
}
