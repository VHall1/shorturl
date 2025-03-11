import { LinkIcon } from "lucide-react";
import Link from "next/link";
import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="bg-background font-sans antialiased">
        <header className="fixed top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
          <div className="h-14 flex items-center container mx-auto px-4 sm:px-6">
            <Link href="/" className="flex items-center">
              <LinkIcon className="size-4 mr-2" />
              <h1 className="text-lg font-bold">ShortURL</h1>
            </Link>
          </div>
        </header>
        <main className="min-h-svh pt-24 container mx-auto px-4 sm:px-6">{children}</main>
      </body>
    </html>
  );
}
