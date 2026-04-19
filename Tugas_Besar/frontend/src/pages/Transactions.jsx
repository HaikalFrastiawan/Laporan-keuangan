import { useState, useEffect } from "react";
import { Plus, Edit2, Trash2, X } from "lucide-react";

export default function Transactions() {
  const [transactions, setTransactions] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const [currentId, setCurrentId] = useState(null);
  const [formData, setFormData] = useState({
    deskripsi: "",
    jumlah: "",
    tanggal: new Date().toISOString().slice(0, 10),
    jenis: "pemasukan",
    catatan: ""
  });

  const fetchTransactions = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/transactions");
      if (res.ok) {
        const data = await res.json();
        setTransactions(data || []);
      }
    } catch (err) {
      console.error("Gagal mengambil data transaksi:", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTransactions();
  }, []);

  const formatRupiah = (number) => {
    return new Intl.NumberFormat("id-ID", { style: "currency", currency: "IDR" }).format(number);
  };

  const handleOpenModal = (tx = null) => {
    if (tx) {
      setIsEdit(true);
      setCurrentId(tx.transactions_id);
      setFormData({
        deskripsi: tx.deskripsi,
        jumlah: tx.jumlah,
        tanggal: tx.tanggal ? tx.tanggal.slice(0,10) : new Date().toISOString().slice(0, 10),
        jenis: tx.jenis,
        catatan: tx.catatan || ""
      });
    } else {
      setIsEdit(false);
      setCurrentId(null);
      setFormData({
        deskripsi: "",
        jumlah: "",
        tanggal: new Date().toISOString().slice(0, 10),
        jenis: "pemasukan",
        catatan: ""
      });
    }
    setShowModal(true);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      ...formData,
      jumlah: parseFloat(formData.jumlah),
      tanggal: new Date(formData.tanggal).toISOString(),
    };

    const url = isEdit 
      ? `http://localhost:8080/api/transactions/${currentId}` 
      : `http://localhost:8080/api/transactions`;
    
    const method = isEdit ? "PUT" : "POST";

    try {
      const res = await fetch(url, {
        method,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });
      if (res.ok) {
        setShowModal(false);
        fetchTransactions();
      }
    } catch (err) {
      console.error("Gagal menyimpan transaksi:", err);
    }
  };

  const handleDelete = async (id) => {
    if(!window.confirm("Apakah Anda yakin ingin menghapus transaksi ini?")) return;
    
    try {
      const res = await fetch(`http://localhost:8080/api/transactions/${id}`, {
        method: "DELETE"
      });
      if (res.ok) {
        fetchTransactions();
      }
    } catch (err) {
      console.error("Gagal menghapus", err);
    }
  };

  if (loading) {
    return <div className="text-neonCyan flex justify-center mt-20 animate-pulse font-mono text-xl">Memuat data transaksi...</div>;
  }

  return (
    <div className="space-y-6 max-w-7xl mx-auto">
      <div className="flex flex-col md:flex-row justify-between items-center bg-darkCard/50 p-6 rounded-2xl border border-white/5 shadow-neon">
        <div>
          <h2 className="text-3xl font-bold text-white font-mono">Manajemen Transaksi</h2>
          <p className="text-textColor mt-2">Kelola pemasukan dan pengeluaran Anda.</p>
        </div>
        <button 
          onClick={() => handleOpenModal()} 
          className="mt-4 md:mt-0 flex items-center gap-2 px-6 py-3 bg-neonCyan/10 text-neonCyan border border-neonCyan rounded-xl hover:bg-neonCyan hover:text-darkBg transition-all font-semibold shadow-neon-hover">
          <Plus size={20} /> Tambah Transaksi
        </button>
      </div>

      <div className="glass rounded-xl overflow-hidden flex flex-col shadow-lg border border-white/5">
        <div className="overflow-x-auto">
          <table className="w-full text-left border-collapse">
            <thead>
              <tr className="bg-white/5 text-textColor text-sm border-b border-white/10 uppercase tracking-wider">
                <th className="p-5 font-medium">Tanggal</th>
                <th className="p-5 font-medium">Deskripsi</th>
                <th className="p-5 font-medium">Jenis</th>
                <th className="p-5 font-medium">Catatan</th>
                <th className="p-5 font-medium text-right">Jumlah</th>
                <th className="p-5 font-medium text-center">Aksi</th>
              </tr>
            </thead>
            <tbody>
              {transactions.map((tx) => (
                <tr key={tx.transactions_id} className="border-b border-white/5 hover:bg-white/5 transition-colors">
                  <td className="p-5 text-sm text-textColor whitespace-nowrap">
                    {new Date(tx.tanggal).toLocaleDateString("id-ID")}
                  </td>
                  <td className="p-5 text-white font-medium">{tx.deskripsi}</td>
                  <td className="p-5">
                    <span className={`text-xs px-3 py-1 rounded-full border ${
                      tx.jenis === "pemasukan" 
                      ? "bg-neonCyan/10 text-neonCyan border-neonCyan/30" 
                      : "bg-neonPurple/10 text-neonPurple border-neonPurple/30"
                    }`}>
                      {tx.jenis.charAt(0).toUpperCase() + tx.jenis.slice(1)}
                    </span>
                  </td>
                  <td className="p-5 text-sm text-textColor truncate max-w-xs">{tx.catatan || "-"}</td>
                  <td className={`p-5 text-right font-bold text-lg font-mono ${tx.jenis === "pemasukan" ? "text-neonCyan" : "text-neonPurple"}`}>
                    {tx.jenis === "pemasukan" ? "+" : "-"}{formatRupiah(tx.jumlah)}
                  </td>
                  <td className="p-5 flex items-center justify-center gap-3">
                    <button 
                      onClick={() => handleOpenModal(tx)}
                      className="p-2 text-textColor hover:text-neonCyan bg-white/5 rounded-lg transition-colors hover:bg-white/10"
                      title="Edit">
                      <Edit2 size={18} />
                    </button>
                    <button 
                      onClick={() => handleDelete(tx.transactions_id)}
                      className="p-2 text-textColor hover:text-red-400 bg-white/5 rounded-lg transition-colors hover:bg-white/10"
                      title="Hapus">
                      <Trash2 size={18} />
                    </button>
                  </td>
                </tr>
              ))}
              {transactions.length === 0 && (
                <tr>
                  <td colSpan={6} className="p-10 text-center text-textColor bg-darkCard/30 font-mono">
                    Belum ada data transaksi. Silakan tambahkan!
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>

      {/* Modal Form */}
      {showModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm p-4">
          <div className="bg-darkCard w-full max-w-lg rounded-2xl border border-white/10 shadow-2xl overflow-hidden relative group">
            <div className="absolute top-0 right-0 -mr-8 -mt-8 w-40 h-40 bg-neonCyan/10 rounded-full blur-3xl"></div>
            
            <div className="px-6 py-5 border-b border-white/10 flex justify-between items-center relative z-10">
              <h3 className="text-xl font-bold text-white tracking-wide">
                {isEdit ? "Edit Transaksi" : "Tambah Transaksi Baru"}
              </h3>
              <button 
                onClick={() => setShowModal(false)}
                className="text-textColor hover:text-white transition bg-white/5 hover:bg-white/10 p-2 rounded-full">
                <X size={20} />
              </button>
            </div>
            
            <form onSubmit={handleSubmit} className="p-6 space-y-5 relative z-10">
              <div>
                <label className="block text-sm text-textColor mb-2 font-medium">Jenis Transaksi</label>
                <select 
                  value={formData.jenis}
                  onChange={(e) => setFormData({...formData, jenis: e.target.value})}
                  className="w-full bg-darkBg border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-neonCyan transition-all"
                  required>
                  <option value="pemasukan">Pemasukan</option>
                  <option value="pengeluaran">Pengeluaran</option>
                </select>
              </div>

              <div>
                <label className="block text-sm text-textColor mb-2 font-medium">Deskripsi</label>
                <input 
                  type="text" 
                  value={formData.deskripsi}
                  onChange={(e) => setFormData({...formData, deskripsi: e.target.value})}
                  className="w-full bg-darkBg border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-neonCyan transition-all placeholder-white/20"
                  placeholder="Contoh: Gaji Bulanan"
                  required 
                />
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="block text-sm text-textColor mb-2 font-medium">Jumlah (Rp)</label>
                  <input 
                    type="number" 
                    min="0"
                    value={formData.jumlah}
                    onChange={(e) => setFormData({...formData, jumlah: e.target.value})}
                    className="w-full bg-darkBg border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-neonCyan transition-all placeholder-white/20"
                    placeholder="50000"
                    required 
                  />
                </div>
                <div>
                  <label className="block text-sm text-textColor mb-2 font-medium">Tanggal</label>
                  <input 
                    type="date" 
                    value={formData.tanggal}
                    onChange={(e) => setFormData({...formData, tanggal: e.target.value})}
                    className="w-full bg-darkBg border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-neonCyan transition-all"
                    required 
                  />
                </div>
              </div>

              <div>
                <label className="block text-sm text-textColor mb-2 font-medium">Catatan (Opsional)</label>
                <textarea 
                  value={formData.catatan}
                  onChange={(e) => setFormData({...formData, catatan: e.target.value})}
                  rows={2}
                  className="w-full bg-darkBg border border-white/10 rounded-xl px-4 py-3 text-white focus:outline-none focus:border-neonCyan transition-all placeholder-white/20"
                  placeholder="Tambahkan catatan jika perlu..."
                ></textarea>
              </div>

              <div className="pt-4 flex justify-end gap-3">
                <button 
                  type="button" 
                  onClick={() => setShowModal(false)}
                  className="px-5 py-2.5 rounded-xl border border-white/10 text-textColor hover:bg-white/5 transition font-medium">
                  Batal
                </button>
                <button 
                  type="submit" 
                  className="px-6 py-2.5 rounded-xl bg-neonCyan text-darkBg border border-neonCyan hover:bg-transparent hover:text-neonCyan transition-all font-bold shadow-neon-hover">
                  Simpan
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}
