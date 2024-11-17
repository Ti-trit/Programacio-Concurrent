--
-- Solució del problema del sopar dels filòsofs
-- amb monitors, segons l'algorisme de Ben-Ari
-- amb l'array amb el nombre de bastonets disponibles
-- us de les Entry families
with Text_IO;
use Text_IO;
with def_monitorPalets;
use def_monitorPalets;

procedure filosofsentry is

   A_MENJAR    : constant Integer := 5;

   p : MonitorPalets;

   task type Filosof is
      entry Start(Idx: in index);
   end Filosof;

   task body Filosof is
      My_Idx : index;

      procedure Pensa is
      begin
         delay Duration(0.05);
      end Pensa;

      procedure Menja is
      begin
         Put_Line(My_Idx'Img & " comensa a menjar");
         delay Duration(0.1);
         Put_Line(My_Idx'Img & " acaba de menjar");
      end Menja;

   begin
      accept Start (Idx : in index) do
         My_Idx := Idx;
      end Start;
      Put_Line ("Filosof " & My_Idx'Img);
      for i in 1..A_MENJAR loop
         Pensa;
         p.Agafa(My_Idx);
         Menja;
         p.Deixa(My_Idx);
      end loop;
      Put_Line("Ja esta ple el filosof " & My_Idx'Img);
      p.Acaba;
      Put_Line("El filòsof " & My_Idx'Img & " se'n va");
   end Filosof;

   type filo is array (index) of Filosof;
   f: filo;

begin
   Put_Line ("Filosofs van a sopar");
   --p.Inicialitza;
   p.Inicialitzar;
   for Idx in index loop
      Put_Line ("Dispara " & Idx'Img);
      f(Idx).Start(Idx);
   end loop;

end filosofsentry;
