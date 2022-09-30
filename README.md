# Obra-blanca

_**v0.0.3b**_

## API
[API Documentation](https://documenter.getpostman.com/view/23460473/2s83mdKj8h)
- [GET] /director/event 👌
  - Devuelve todos los eventos guardados en el sistema
- [POST] /director/event 👌
  - Crea un nuevo evento y devuelve su Id
- [GET] /director/event?id={UUID} 👌
  - Devuelve únicamente un evento de Id `UUID` guardado en el sistema (si existe)
- [GET] /director/product?event={UUID} 👌
  - Devuelve todos los productos del evento de Id `UUID`
- [POST] /director/product?event={UUID} 👌
  - Agrega un nuevo producto al evento de Id`UUID`
- [GET] /director/inventory?event={eUUID}&product={pUUID} 👌
  - Obtiene los datos de inventario del producto de Id `pUUID` en el evento `eUUID`
- [POST] /director/inventory?event={eUUID}&product={pUUID}&amount={amt} 👌
  - Agrega un nuevo inventario inicial de cantidad `amt` (creando un registro de existencias) para un producto 
de Id `pUUID` en el evento `eUUID`
- [POST] /director/promoter 👌
  - Agrega un promotor al catálogo de promotores 
- [GET] /director/promoter 👌
  - Obtener el catálogo de promotores
- [DELETE] /director/promoter?name={name} 👌
  - Elimina a un promotor de nombre `name` del catálogo de promotores
- [POST] /director/distributor 👌
  - Agrega un promotor al catálogo de distribuidores
- [GET] /director/distributor 👌
  - Obtener el catálogo de distribuidores
- [DELETE] /director/distributor?name={name} 👌
  - Elimina a un distribuidor de nombre `name` del catálogo de promotores
- [POST] /director/assignment/event={UUID}
  - Crea un nuevo contrato de negociación para un distribuidor, asociado a un evento dado
- [GET] /promoter/product?event={UUID} 👌
  - Devuelve todos los productos del evento de Id`UUID`
- [GET] /promoter/assignment/event={name} 👌
  - Devuelve todos los contratos de negociación por distribuidor para el evento de nombre `name`
- [POST] /promoter/assignment?event={eName}&distributor={pName} 👌
  - Crea un nuevo contrato de negociación para el distribuidor `eName` en el evento `pName`
- [GET] /promoter/assignment/event={eName}&distributor={pName} 👌
  - Devuelve el detalle del contrato de negociación del distribuidor de nombre `name`
- [PUT] /promoter/assignment?event={eName}&distributor={name}
  - Añade o actualiza un producto al contrato de negociación del distribuidor de nombre `name`
- [DELETE] /promoter/assignment/product?id={UUID}
  - Quita un producto del contrato de negociación de Id `UUID`
- [GET] /promoter/inventory/product={UUID} 👌
  - Devuelve el inventario del producto de Id `UUID`


- ~~/promoter/set-product-aside~~

## Casos de uso
- [x] Yo como usuario quiero tener acceso seguro a la aplicación
- [x] Yo como Administrador quiero registrar inventarios de nuevos productos
para permitir a los promotores apartarlo para los distribuidores interesados 
a partir de un evento de Obra Blanca
- [x] Yo como Administrador quiero abrir y cerrar ciclos por medio de eventos 
para llevar el seguimiento de los nuevos productos de Lamosa a través del tiempo
- [x] Yo como Administrador quiero registrar promotores para controlar el acceso 
al sistema
- [ ] Yo como Promotor quiero tener la lista de nuevos productos para poder llevar 
el seguimiento de los distribuidores que están interesados en venderlo
- [ ] Yo como Promotor quiero conocer la cantidad de inventario restante de 
los nuevos productos de Lamosa para poder promoverlo y venderlo a los distribuidores
- [x] Yo como Administrador quiero asignar un distribuidor a un promotor para llevar
el seguimiento de ese proceso de negociación
- [ ] Yo como Administrador quiero procesar la información transaccional 
generada por el sistema para generar reportes que permitan tomar mejores 
decisiones relativas a nuevos productos en Lamosa

## Dominio
- Distributor (distributor)
- Promotor (promoter)
- Evento (event)
  - Negociación (negotiation)
  - Producto (product)
    - Inventario (inventory)
