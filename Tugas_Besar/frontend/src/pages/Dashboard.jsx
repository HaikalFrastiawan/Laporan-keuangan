import { useEffect, useState } from "react";
import { ArrowUpRight, ArrowDownRight, Wallet } from "lucide-react";

export default function Dashboard() {
  // Using Mock Data structure as fallback if backend is not yet fully linked
  const initialData = {
    total_pemasukan: 0,
    total_pengeluaran: 0,
    saldo_akhir: 0,
    recent_transactions: []
  };

  const [data, setData] = useState(initialData);
  const [loading, setLoading] = useState(true);

  // Example frontend fetch call to our Go backend
  useEffect(() => {
    const fetchDashboard = async () => {
      try {
        const response = await fetch("http://localhost:8080/api/dashboard");
        if (response.ok) {
          const result = await response.json();
          setData(result);
        }
      } catch (err) {
        console.warn("Backend belum berjalan. Menggunakan data statis.");
        // Mock fallback if DB is not running
        setData({
          total_pemasukan: 15500000,
          total_pengeluaran: 4200000,
          saldo_akhir: 11300000,
          recent_transactions: [
            { transactions_id: 1, deskripsi: "Gaji Bulan Ini", jumlah: 15000000, tanggal: "2026-04-10T10:00:00Z", jenis: "pemasukan" },
            { transactions_id: 2, deskripsi: "Bayar Kost", jumlah: 1500000, tanggal: "2026-04-12T12:00:00Z", jenis: "pengeluaran" },
            { transactions_id: 3, deskripsi: "Makan Siang", jumlah: 50000, tanggal: "2026-04-15T13:00:00Z", jenis: "pengeluaran" }
          ]
        });
      } finally {
        setLoading(false);
      }
    };
    fetchDashboard();
  }, []);

  const formatRupiah = (number) => {
    return new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(number);
  };

  if (loading) {
    return <div className="text-neonCyan flex justify-center mt-20 animate-pulse">Memuat data...</div>;
  }

  return (
    <div className="space-y-6 max-w-7xl mx-auto">
      
      {/* 1. Final Balance Section Highlighted */}
      <div className="glass rounded-2xl p-8 relative overflow-hidden group shadow-neon transition-shadow">
        <div className="absolute top-0 right-0 -mr-8 -mt-8 w-40 h-40 bg-neonPurple/20 rounded-full blur-3xl group-hover:bg-neonPurple/30 transition-colors"></div>
        <div className="relative z-10 flex flex-col md:flex-row items-center justify-between">
          <div>
            <p className="text-textColor font-medium tracking-wider uppercase flex items-center gap-2">
              <Wallet size={18} className="text-neonCyan" />
              Saldo Akhir Anda
            </p>
            <h2 className="text-4xl md:text-5xl font-bold text-white mt-2 font-mono">
              {formatRupiah(data.saldo_akhir)}
            </h2>
          </div>
          <button className="mt-4 md:mt-0 px-6 py-3 bg-neonCyan/10 text-neonCyan border border-neonCyan rounded-xl hover:bg-neonCyan hover:text-darkBg transition-all font-semibold shadow-neon-hover">
            + Tambah Transaksi
          </button>
        </div>
      </div>

      {/* 2. Grid Cards (Income vs Expense) */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="glass rounded-xl p-6 border-l-4 border-l-neonCyan">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-textColor text-sm">Total Pemasukan</p>
              <h3 className="text-2xl font-bold text-white mt-1">{formatRupiah(data.total_pemasukan)}</h3>
            </div>
            <div className="w-10 h-10 rounded-full bg-neonCyan/20 flex items-center justify-center">
              <ArrowUpRight className="text-neonCyan" />
            </div>
          </div>
        </div>

        <div className="glass rounded-xl p-6 border-l-4 border-l-neonPurple">
          <div className="flex justify-between items-start">
            <div>
              <p className="text-textColor text-sm">Total Pengeluaran</p>
              <h3 className="text-2xl font-bold text-white mt-1">{formatRupiah(data.total_pengeluaran)}</h3>
            </div>
            <div className="w-10 h-10 rounded-full bg-neonPurple/20 flex items-center justify-center">
              <ArrowDownRight className="text-neonPurple" />
            </div>
          </div>
        </div>
      </div>

      {/* 3. Recent Transactions Table */}
      <div className="glass rounded-xl overflow-hidden flex flex-col">
        <div className="p-6 border-b border-white/5">
          <h3 className="text-lg font-semibold text-white">Transaksi Terbaru</h3>
        </div>
        
        <div className="overflow-x-auto">
          <table className="w-full text-left border-collapse">
            <thead>
              <tr className="bg-white/5 text-textColor text-sm">
                <th className="p-4 font-medium">Tanggal</th>
                <th className="p-4 font-medium">Deskripsi</th>
                <th className="p-4 font-medium">Jenis</th>
                <th className="p-4 font-medium text-right">Jumlah</th>
              </tr>
            </thead>
            <tbody>
              {data.recent_transactions.map((tx, idx) => (
                <tr key={idx} className="border-b border-white/5 hover:bg-white/5 transition-colors">
                  <td className="p-4 text-sm text-textColor whitespace-nowrap">
                    {new Date(tx.tanggal).toLocaleDateString("id-ID")}
                  </td>
                  <td className="p-4 text-white font-medium">{tx.deskripsi}</td>
                  <td className="p-4">
                    <span className={`text-xs px-2 py-1 rounded-full border ${
                      tx.jenis === "pemasukan" 
                      ? "bg-neonCyan/10 text-neonCyan border-neonCyan/30" 
                      : "bg-neonPurple/10 text-neonPurple border-neonPurple/30"
                    }`}>
                      {tx.jenis.charAt(0).toUpperCase() + tx.jenis.slice(1)}
                    </span>
                  </td>
                  <td className={`p-4 text-right font-medium ${tx.jenis === "pemasukan" ? "text-neonCyan" : "text-neonPurple"}`}>
                    {tx.jenis === "pemasukan" ? "+" : "-"}{formatRupiah(tx.jumlah)}
                  </td>
                </tr>
              ))}
              {data.recent_transactions.length === 0 && (
                <tr>
                  <td colSpan={4} className="p-6 text-center text-textColor">
                    Belum ada aktivitas transaksi
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
