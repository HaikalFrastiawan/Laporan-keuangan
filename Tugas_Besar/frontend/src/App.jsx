import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Sidebar from "./components/layout/Sidebar";
import Header from "./components/layout/Header";
import Dashboard from "./pages/Dashboard";

function App() {
  return (
    <Router>
      <div className="flex h-screen bg-darkBg overflow-hidden">
        <Sidebar />
        
        <div className="flex-1 flex flex-col overflow-hidden">
          <Header />
          
          <main className="flex-1 overflow-x-hidden overflow-y-auto bg-darkBg p-6">
            <Routes>
              {/* Default Redirect to Dashboard */}
              <Route path="/" element={<Navigate to="/dashboard" replace />} />
              <Route path="/dashboard" element={<Dashboard />} />
              {/* Other pages placeholder */}
              <Route path="/transactions" element={<div className="text-textColor">Halaman Transaksi</div>} />
              <Route path="/categories" element={<div className="text-textColor">Halaman Kategori</div>} />
              <Route path="/reports" element={<div className="text-textColor">Halaman Laporan</div>} />
              <Route path="/profile" element={<div className="text-textColor">Halaman Profil</div>} />
            </Routes>
          </main>
        </div>
      </div>
    </Router>
  );
}

export default App;
