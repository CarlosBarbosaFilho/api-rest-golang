create table personas (
    id serial primary key,
    name varchar 255,
    history varchar 255
)

insert into personas (name, history) values ('Carlos Barbosa G. Filho', 'Profissional de TI, nascido em 03/12/1976, casado com 3 filhas');
insert into personas (name, history) values ('Barbara de Assis', 'Profissional de TI, nascida em 01/02/1985, casada com 3 filhas');
insert into personas (name, history) values ('Gabryella Waleska', 'Profissional de TI, nascido em 02/02/1995, casado com 3 filhas');
insert into personas (name, history) values ('Maria Alice', 'Filha nascida em 22/03/2022, filha do meio');
insert into personas (name, history) values ('Maria Helena', 'Filha nascida em 29/12/2023, filha mais nova');
