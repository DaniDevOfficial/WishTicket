import { Outlet, ScrollRestoration } from "react-router-dom";



export function DefaultLayout() {
    return (
        <div className="w-full">
            <div className="min-h-screen w-full">
                <main className="mb-8">
                    <div className="max-w-5xl mx-auto">
                        <Outlet />
                    </div>
                </main>
            </div>

            <ScrollRestoration />
        </div>
    );
}