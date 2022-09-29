package ortalamahesaplama

// Vize Notu 0 ile 100 arasında olabilir
// Final Notu 0 ile 100 arasında olabilir
// İnsiyatif 0 ile 1 arasında olabilir
// Ortalama 0 ile 100 arasında olabilir
// Ortalama 65'ten küçükse ve ortalama*(1+insiyatif) 65'ten büyükse ortalama 65 olmalı
// Ortalama 65'ten büyükse ortalama değişmemelidir
// Çıktı olarak ortalama dönmelidir
// OrtalamaHesapla fonksiyonu `go test ./...` komutu ile test edilmelidir
// Testlerin hepsi PASS olmalıdır

func OrtalamaHesapla(vizeNot int, finalNot int, insiyatif float64) float64 {
	ortalama := (float64(vizeNot) * 0.4) + (float64(finalNot) * 0.6)

	if ortalama < 65 && ortalama * (1 + insiyatif) > 65 {
		ortalama = 65
	}

	return ortalama
}

/*

Aşağıdaki örnekte vize, final, ve insiyatif notlarının minimum ve maksimum değerleri de kontrol ediliyor. Üstte verilen  koşula uymadıkları taktirde sistem hata verir.

func OrtalamaHesapla(vizeNot int, finalNot int, insiyatif float64) (float64 error) {
	ortalama := (float64(vizeNot) * 0.4) + (float64(finalNot) * 0.6)
	if vizeNot < 0 || vizeNot > 100 {
        return errors.New("Geçersiz vize notu!")
    }

	if finalNot < 0 || finalNot > 100 {
        return errors.New("Geçersiz final notu!")
    }

	if insiyatif < 0 || insiyatif > 1 {
        return errors.New("Geçersiz insiyatif!")
    }

	if ortalama < 65 && ortalama * (1 + insiyatif) > 65 {
		ortalama = 65
	}

	return ortalama
}
*/