using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace RuGa.Models.Sections_Models;

public class Section
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int id_section { get; set; }
    public string   SectionNameModel { get; set; }  =   string.Empty;
    public string?   SectionDescriptionModel { get; set; }
    public  string?  id_user_fk{ get; set; }
}
