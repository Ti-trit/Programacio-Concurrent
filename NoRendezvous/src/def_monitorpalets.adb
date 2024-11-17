with Ada.Text_IO;
use Ada.Text_IO;

package body def_monitorpalets is

   acabats : Integer := 0;

     function esq(i: index) return index is
       begin
         return index((Integer(i) + NUMFILOSOFS - 1) mod NUMFILOSOFS);
     end esq;

     function dre(i: index) return index is
       begin
         return index((Integer(i) + 1) mod NUMFILOSOFS);
   end dre;

   protected body MonitorPalets is

      function get_palets(Idx: in index) return Integer is
      begin
         return palets(Idx);
         end get_palets;

    entry Agafa (for Idx in index) when palets(Idx)=2  is

    begin
            --agafa els bastonets dels veinats
          palets(esq(Idx)) := palets(esq(Idx)) - 1;
          palets(dre(Idx)) := palets(dre(Idx)) - 1;
 end Agafa;

 procedure Deixa (Idx : in index)  is
    begin
            --retorna els bastonets als veinats
          palets(esq(Idx)) := palets(esq(Idx)) + 1;
          palets(dre(Idx)) := palets(dre(Idx)) + 1;
      -- end if;
 end Deixa;

 procedure Inicialitzar  is
    begin
       for i in index loop
         palets(i) := 2;
    end loop;
 end Inicialitzar;

    procedure Acaba is
       begin
    acabats:= acabats+1;
    end Acaba;

 end MonitorPalets;

  end def_monitorpalets;
--     task body MonitorPalets is
--        type tpalets is array (index) of integer;
--        palets : tpalets;
--     begin
--        accept Inicialitza do
--           for i in index loop
--             palets(i) := 2;
--           end loop;
--        end Inicialitza;
--        loop
--           for I in index loop
--              select
--                 when palets(I) = 2 =>
--                    accept Agafa (I) (Idx : in index) do
--                       palets(esq(Idx)) := palets(esq(Idx)) - 1;
--                       palets(dre(Idx)) := palets(dre(Idx)) - 1;
--                    end Agafa;
--              or
--                 accept Deixa (Idx : in index) do
--                    palets(esq(Idx)) := palets(esq(Idx)) + 1;
--                    palets(dre(Idx)) := palets(dre(Idx)) + 1;
--                 end Deixa;
--              or
--                 accept Acaba  do
--                    acabats := acabats + 1;
--                 end Acaba;
--                 exit;
--              else
--                 null; -- Necessari per evitar que la tasca quedi suspesa quan no hi ha cridades
--              end select;
--           end loop;
--           exit when acabats = NUMFILOSOFS;
--        end loop;
--
--        end MonitorPalets;
--
--  end def_monitorpalets;
