package model

type Parfume struct {
	IdParfume uint `gorm:"primaryKey;colum:id_parfume" json:"id_parfume"`
	NamaParfume string `gorm:"colum:nama_parfume" json:"nama_parfume"`
	JenisParfume string `gorm:"colum:jenis_parfume" json:"jenis_parfume"`
	Deskripsi string `gorm:"colum:deskripsi" json:"deskripsi"`
	TahunPeluncuran uint `gorm:"colum:tahun_peluncuran" json:"tahun_peluncuran"`
	Harga uint `gorm:"colum:harga" json:"harga"`
	CreatedAt string `gorm:"colum:created_at" json:"created_at"`
	UpdatedAt string `gorm:"colum:updated_at" json:"updated_at"`
}