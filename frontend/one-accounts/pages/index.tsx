import React from 'react'
import ButtonAppBar from '../component/ButtonAppBar'
import AccountDetails from '../component/AccountDetails'
import AccountDetailsTabs from  '../component/AccountDetailsTabs'
import AccountDetailsOperationNavi from  '../component/AccountDetailsOperationNavi'

const Home = () => {
    return (
        <>
        <ButtonAppBar />
        <AccountDetailsTabs />
        <AccountDetails />
        <AccountDetailsOperationNavi />
        </>
    )
}

export default Home;