using System;

namespace RuGa.DTOs.Sections_DTOs;

public class SectionDTO
{
    public int id_section { get; set; }
    public string   SectionNameDto { get; set; }    =   string.Empty;
    public string?  SectionDescriptionDto { get; set; }
    public  string?  id_user_fk { get; set; }
}