using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Identity;
using Microsoft.IdentityModel.Tokens;

using RuGa.Data.DataBases;
using RuGa.Data.Identities;
using RuGaPasswordManager.Services;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using RuGaPasswordManager.Interfaces.Account_Interfaces;

namespace RuGa;

public class Program
{
    public static void Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        // Add services to the container.
        // =============================================================================================
        builder.Services.AddControllers();
        // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
        builder.Services.AddEndpointsApiExplorer();

        // =============================================================================================
        // ======================================== DBCONTEXT ==========================================
        var conn    =   builder.Configuration.GetConnectionString("PostrgresqlConnection");
        builder.Services.AddDbContext<ApiDbContext>(options =>  options.UseNpgsql(conn));
        builder.Services.AddSwaggerGen();

        // =============================================================================================
        // ========================================= IDENTITY ==========================================
        // builder.Services.AddAuthorization();
        /* DEFINICION DE IDENTITY (USUARIO Y ROLES). */
        builder.Services.AddIdentity<AppUser,   IdentityRole>(options   =>  {
            options.Password.RequireDigit = true;
            options.Password.RequiredLength = 5;
            options.Password.RequireUppercase = true;
            options.Password.RequireLowercase   = true;
            options.Password.RequireNonAlphanumeric = true;
        })
        .AddEntityFrameworkStores<ApiDbContext>();// ALMACENAMIENTO EN LA BASE DE DATOS.

        /* DEFINICION DE ESQUEMAS PARA LA AUTENTICACION. SE USARAN TOKENS *JWT-BEARER*. */
        builder.Services.AddAuthentication(options =>   {
            options.DefaultAuthenticateScheme   =
            options.DefaultChallengeScheme      =
            options.DefaultForbidScheme         =
            options.DefaultScheme               =
            options.DefaultSignInScheme         =
            options.DefaultSignOutScheme        =   JwtBearerDefaults.AuthenticationScheme;
        })
        .AddJwtBearer(options   =>  {// DEFINICION DE LOS TOKENS JWT.
            options.TokenValidationParameters   =   new TokenValidationParameters   {
                ValidateIssuer  =   true,
                ValidIssuer     =   builder.Configuration["JWT:Issuer"],

                ValidateAudience=   true,
                ValidAudience   =   builder.Configuration["JWT:Audience"],

                ValidateIssuerSigningKey    =   true,
                IssuerSigningKey    =   new SymmetricSecurityKey(
                    System.Text.Encoding.UTF8.GetBytes(builder.Configuration["JWT:SignInKey"])
                )
            };
        });


        // =============================================================================================
        // ========================================= SERVICIOS =========================================
        builder.Services.AddScoped<ITokenService,   TokenService>();



        var app = builder.Build();

        // Configure the HTTP request pipeline.
        if (app.Environment.IsDevelopment())
        {
            app.UseSwagger();
            app.UseSwaggerUI();
        }

        app.UseHttpsRedirection();

        app.MapControllers();

        app.UseAuthentication();//MIDLEWARE DE AUTENTICACION.
        app.UseAuthorization();//MIDLEWARE DE AUTORIZACION.

        app.Run();
    }
}
