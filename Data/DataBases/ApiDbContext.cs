using System;
using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Identity.EntityFrameworkCore;
using Microsoft.AspNetCore.Identity;

using RuGa.Data.Identities;
using RuGa.Models.Sections_Models;

namespace RuGa.Data.DataBases;

public class ApiDbContext   :   IdentityDbContext<AppUser>
{
    public ApiDbContext(DbContextOptions<ApiDbContext> options) : base(options){ 
        AppContext.SetSwitch("Npgsql.EnableLegacyTimestampBehavior", true);
    }

    // DBSET PARA LA TABLA DE SECCION.
    public  DbSet<Section> Section { get; set; }

    //IDENTITY ROLES.
    // PARA CREAR USUARIOS, ES NECESARIO TENER LOS ROLES A LO QUE SERAN ASIGNADOS.
    protected override void OnModelCreating(ModelBuilder    builder){
        base.OnModelCreating(builder);

        //CREACION DE ROLES.
        List<IdentityRole>  roles = new List<IdentityRole>{
            new IdentityRole{
                Name            =   "Admin ",
                NormalizedName   =   "ADMIN"
            },
            new IdentityRole    {
                Name            =   "User",
                NormalizedName  =   "USER"
            }
        };

        // AÑADIR LOS ROLES AL BUILDER.
        builder.Entity<IdentityRole>().HasData(roles);
    }
}