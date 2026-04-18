import { Link, useLocation } from "react-router-dom";
import { LayoutDashboard, ReceiptText, Tags, PieChart, Users, LogOut } from "lucide-react";

const Sidebar = () => {
  const location = useLocation();

  const menuItems = [
    { name: "Dashboard", path: "/dashboard", icon: <LayoutDashboard size={20} /> },
    { name: "Manajemen Transaksi", path: "/transactions", icon: <ReceiptText size={20} /> },
    { name: "Manajemen Kategori", path: "/categories", icon: <Tags size={20} /> },
    { name: "Laporan Keuangan", path: "/reports", icon: <PieChart size={20} /> },
    { name: "Manajemen User", path: "/profile", icon: <Users size={20} /> },
  ];

  return (
    <aside className="w-64 glass flex flex-col h-full border-r border-neonTeal/20">
      {/* Brand */}
      <div className="h-16 flex items-center justify-center border-b border-neonTeal/20">
        <h1 className="text-xl font-bold text-neonCyan tracking-wider flex items-center gap-2">
          <div className="w-8 h-8 rounded-full bg-neonCyan/20 border border-neonCyan flex items-center justify-center shadow-neon">
            <span className="text-neonCyan">LK</span>
          </div>
          Vibe Finance
        </h1>
      </div>

      {/* Navigation */}
      <nav className="flex-1 overflow-y-auto py-6 px-4 space-y-2">
        {menuItems.map((item) => {
          const isActive = location.pathname.startsWith(item.path);
          return (
            <Link
              key={item.name}
              to={item.path}
              className={`flex items-center gap-3 px-4 py-3 rounded-lg transition-all duration-300 ${
                isActive
                  ? "bg-neonTeal/20 text-neonCyan shadow-neon border border-neonTeal/50"
                  : "text-textColor hover:bg-darkCard hover:text-white"
              }`}
            >
              {item.icon}
              <span className="font-medium">{item.name}</span>
            </Link>
          );
        })}
      </nav>

      {/* Logout Bottom */}
      <div className="p-4 border-t border-neonTeal/20">
        <button className="flex items-center gap-3 w-full px-4 py-3 rounded-lg text-red-400 hover:bg-red-500/10 hover:text-red-300 transition-colors">
          <LogOut size={20} />
          <span className="font-medium">Logout</span>
        </button>
      </div>
    </aside>
  );
};

export default Sidebar;
