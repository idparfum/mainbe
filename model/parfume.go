package model

type Parfume struct {
	IdParfume       uint   `gorm:"primaryKey;column:id_parfume" json:"id_parfume"`
	NamaParfume     string `gorm:"column:nama_parfume" json:"nama_parfume"`
	JenisParfume    string `gorm:"column:jenis_parfume" json:"jenis_parfume"`
	Deskripsi       string `gorm:"column:deskripsi" json:"deskripsi"`
	TahunPeluncuran uint   `gorm:"column:tahun_peluncuran" json:"tahun_peluncuran"`
	Harga           uint   `gorm:"column:harga" json:"harga"`
	CreatedAt       string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       string `gorm:"column:updated_at" json:"updated_at"`
}
