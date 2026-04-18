import { Bell, Search } from "lucide-react";

const Header = () => {
  return (
    <header className="h-16 glass border-b border-neonTeal/20 flex items-center justify-between px-6 z-10 w-full">
      {/* Left side info */}
      <div className="flex-1">
        {/* Can put breadcrumbs or search here */}
        <div className="relative w-64">
          <span className="absolute inset-y-0 left-0 pl-3 flex items-center text-textColor">
            <Search size={18} />
          </span>
          <input
            type="text"
            className="w-full bg-darkCard text-textColor rounded-full py-1.5 pl-10 pr-4 focus:outline-none focus:ring-1 focus:ring-neonCyan border border-transparent focus:border-neonCyan transition-all"
            placeholder="Cari transaksi..."
          />
        </div>
      </div>

      {/* Right side Profile & Notification */}
      <div className="flex items-center gap-6">
        <button className="text-textColor hover:text-neonCyan transition-colors relative">
          <Bell size={20} />
          <span className="absolute -top-1 -right-1 flex h-3 w-3">
            <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-neonPurple opacity-75"></span>
            <span className="relative inline-flex rounded-full h-3 w-3 bg-neonPurple"></span>
          </span>
        </button>
        
        <div className="flex items-center gap-3 border-l border-white/10 pl-6 cursor-pointer group">
          <div className="text-right hidden md:block">
            <p className="text-sm font-semibold text-white group-hover:text-neonCyan transition-colors">Haikal Frastiawan</p>
            <p className="text-xs text-textColor">Administrator</p>
          </div>
          <div className="w-10 h-10 rounded-full bg-gradient-to-tr from-neonCyan to-neonPurple p-0.5 shadow-neon">
            <img
              src="https://ui-avatars.com/api/?name=Haikal+F&background=1F2833&color=66FCF1"
              alt="Profile"
              className="w-full h-full rounded-full object-cover border-2 border-darkBg"
            />
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
