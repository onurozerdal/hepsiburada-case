## Hepsiburada Search & Navigation Data Team Assignment

Bu case içerisinde;

* Kullanıcıların history bilgisinin bir json üzerinden okunması,
* Bu bilgilerden kişilerin ilgi alanlarının tespit edilmesi,
* Edilen ilgi alanlarına göre kişiye o ilgi alanına ait bestseller ürünlerin önerilmesi,
* Eğer kişinin bir history bilgisi yoksa genel bestseller ürünler önerilmesi,
* user-id ve product-id ile history bilgisinin silinebilmesi

işlemleri uygulanmıştır.

Ana dizinde gözüken modüllerin işlevleri şu şekildedir;

* api: 8080 portunda çalışan ve öneri sorgularının yapıldığı uygulamadır.
* etl-process: 
    * Bestseller ürünler hesaplanarak recommendation engine tablolarına yazılır.
    * Product - Category eşleşmeleri recommendation tablolarına yazılır.
* init: Recommendation engine tarafı için uygun database ve tablolarının oluşturulması için çalıştırılır.
* view-producer: Browsing history bilgilerini jsondan alarak kafkaya atar. 
* stream-reader: Kafkada history topic bilgisini dinler ve buradan gelen değerleri recommendation engine tarafındaki tablolara yazar.

Projeyi çalıştırmak için docker-compose.yml hazırladım. Ama compose dosyası veya docker image larımda bir sorun var ve burayı düzeltmeye zamanım yetmedi.
Localde çalıştırılması durumunda ilk çalıştırılan modül init olmalıdır. Bu sayede database ve tablolar oluşacaktır. 
Yedek olması açısından yaratılması gereken database ve table alter bilgilerini proje içerisindeki init.sql içerisine koydum.

Api modülü 8080 portu üzerinden çalışır. browsingHistories (GET), bestsellerProducts (GET) ve deleteHistory (DELETE) endpointleri bu servis üzerinden erişilebilir.
Requestlerin detaylarına proje içerisindeki postman collectiondan erişilebilir. 
