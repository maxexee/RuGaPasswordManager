using System;
using Microsoft.AspNetCore.Mvc;

using RuGa.Data.DataBases;
using RuGa.DTOs.Sections_DTOs;
using RuGa.Models.Sections_Models;

namespace RuGa.Controllers.Sections
{
    [Route("user/[controller]")]
    [ApiController]
    public class SectionController : ControllerBase
    {
        private readonly ApiDbContext   _context;
        public SectionController(ApiDbContext    context)
        {
            _context    = context;
        }

        // [Authorize]
        [HttpPost]
        public async Task<ActionResult<SectionDTO>> Post(SectionPostDTO   SectionPostDTO){
            // CREACION DE LA NUEVA SECCION EN BASE A *SectionPostDTO*.
            var sectionNew    =   new Section(){
                SectionNameModel         =  SectionPostDTO.SectionNameDtoPost,
                SectionDescriptionModel  =  SectionPostDTO.SectionDescriptionDtoPost,
            };

            // GUARDADO DE LA NUEVA SECCION EN LA BASE DE DATOS.
            await   _context.Section.AddAsync(sectionNew);
            await   _context.SaveChangesAsync();

            // RETORNO QUE MUESTRA A LA NUEVA SECCION.
            var sectionDto =   new SectionDTO(){
                id_section              =   sectionNew.id_section,
                SectionNameDto          =   sectionNew.SectionNameModel,
                SectionDescriptionDto   =   sectionNew.SectionDescriptionModel
            };

            return Ok();
        }
    }
}
