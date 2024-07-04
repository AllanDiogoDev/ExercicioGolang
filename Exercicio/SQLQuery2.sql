-- Allan Diogo Bastos 03/07/2024
-- criando a view

CREATE VIEW vw_TotalHorasPorDiaSemana AS
SELECT
    c.CadastroId AS CadastroId,
    ds.Descricao AS DiaSemana,
    COALESCE(SUM(ht.Horas), 0) AS TotalHoras
FROM
    transdata.dbo.DiasSemana ds
CROSS JOIN 
    (SELECT DISTINCT CadastroId FROM transdata.dbo.HorariosTrabalho) c
LEFT JOIN 
    transdata.dbo.HorariosTrabalho ht 
    ON c.CadastroId = ht.CadastroId
    AND ds.DiaId = ht.DiaId
GROUP BY
    c.CadastroId,
    ds.Descricao,
    ds.DiaId;

-- consultando a view

SELECT *
FROM vw_TotalHorasPorDiaSemana
ORDER BY CadastroId,
    CASE DiaSemana
        WHEN 'Segunda' THEN 2
        WHEN 'Terça' THEN 3
        WHEN 'Quarta' THEN 4
        WHEN 'Quinta' THEN 5
        WHEN 'Sexta' THEN 6
        WHEN 'Sábado' THEN 7
        WHEN 'Domingo' THEN 1
        ELSE 8
    END;


