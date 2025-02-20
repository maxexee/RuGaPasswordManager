using System;
using RuGa.Data.Identities;

namespace RuGaPasswordManager.Interfaces.Account_Interfaces
{
    public interface ITokenService
    {
        string  CreateToken(AppUser user);
    }
}